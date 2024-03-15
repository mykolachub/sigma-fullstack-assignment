import React, { ReactNode } from 'react';
import { Route, Routes } from 'react-router-dom';
import Home from './views/home/Home';
import CreateUser from './views/create/CreateUser';
import './App.css';
import UpdateUser from './views/update/Update';
import Signup from './views/signup/Signup';

interface PageLayoutProps {
  children: ReactNode;
}

const PageLayout = ({ children }: PageLayoutProps) => (
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
        <Route
          path="/signup"
          element={
            <PageLayout>
              <Signup />
            </PageLayout>
          }
        />
      </Routes>
    </div>
  );
}

export default App;
