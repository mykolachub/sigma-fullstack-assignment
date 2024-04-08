package aerospike

import (
	"errors"
	"sigma-test/internal/entity"

	"github.com/adrianbrad/queue"
	aero "github.com/aerospike/aerospike-client-go/v7"
)

type AerospikePageConfig struct {
	Namespace string
	Set       string
}

type PageRepo struct {
	client *aero.Client
	cfg    AerospikePageConfig
}

func NewPageRepo(client *aero.Client, cfg AerospikePageConfig) *PageRepo {
	return &PageRepo{
		client: client,
		cfg:    cfg,
	}
}

func (r *PageRepo) GetPage(name string) (entity.Page, error) {
	key, err := aero.NewKey(r.cfg.Namespace, r.cfg.Set, name)
	if err != nil {
		return entity.Page{}, err
	}

	var rec *aero.Record
	rec, err = r.client.Get(nil, key, "count")
	if err != nil {
		return entity.Page{}, errors.New("no such key")
	}
	count := rec.Bins["count"].(int)
	return entity.Page{Name: name, Count: count}, nil
}

func (r *PageRepo) ResetPageCount(name string) error {
	key, err := aero.NewKey(r.cfg.Namespace, r.cfg.Set, name)
	if err != nil {
		return err
	}

	updateBin := aero.NewBin("count", 0)
	policy := aero.NewWritePolicy(0, 0)
	err = r.client.PutBins(policy, key, updateBin)
	if err != nil {
		return errors.New("failed to reset count")
	}
	return nil
}

func (r *PageRepo) TrackPage(name string) error {
	key, err := aero.NewKey(r.cfg.Namespace, r.cfg.Set, name)
	if err != nil {
		return err
	}

	policy := aero.NewWritePolicy(0, 0)
	incrementBy := aero.NewBin("count", 1)
	_, err = r.client.Operate(policy, key, aero.AddOp(incrementBy))
	return err
}

func (r *PageRepo) BatchTrackPages(q *queue.Linked[string]) error {
	var batchRecords []aero.BatchRecordIfc
	for v := range q.Iterator() {
		key, err := aero.NewKey(r.cfg.Namespace, r.cfg.Set, v)
		if err != nil {
			return err
		}

		incrementBy := aero.NewBin("count", 1)
		record := aero.NewBatchWrite(aero.NewBatchWritePolicy(), key, aero.AddOp(incrementBy))
		batchRecords = append(batchRecords, record)
	}
	return r.client.BatchOperate(aero.NewWriteBatchPolicy(), batchRecords)
}
