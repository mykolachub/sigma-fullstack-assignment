import { Route, Routes } from 'react-router-dom';
import Home from './views/home/Home';
import CreateUser from './views/create/CreateUser';

import './App.css';
import UpdateUser from './views/update/Update';

const PageLayout = ({ children }) => (
  <div className="app__wrapper h-screen">{children}</div>
);

function App() {
  return (
    <div className="app__container">
      <Routes>
        <Route
          path="/"
          element={
            <PageLayout>
              <Home />
            </PageLayout>
          }
        />
        <Route
          path="/create"
          element={
            <PageLayout>
              <CreateUser />
            </PageLayout>
          }
        />
        <Route
          path="/update"
          element={
            <PageLayout>
              <UpdateUser />
            </PageLayout>
          }
        />
      </Routes>
    </div>
  );
}

export default App;
