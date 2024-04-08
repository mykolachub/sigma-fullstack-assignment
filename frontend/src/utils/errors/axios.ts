import axios from 'axios';
import { ServerErrorResponse } from '../../types/server';
import config from '../config';

const UNKNOWN_ERROR = 'unknown server error';
const { service_code } = config;

export const handleAxiosError = (error: unknown): string => {
  if (axios.isAxiosError(error)) {
    const data = error.response?.data as ServerErrorResponse;
    const message = service_code[data.code] || data.message || UNKNOWN_ERROR;
    return message;
  }
  return UNKNOWN_ERROR;
};
