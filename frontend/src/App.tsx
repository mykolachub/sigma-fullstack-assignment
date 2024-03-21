import React from 'react';
import { Route, Routes } from 'react-router-dom';
import './App.css';

import Home from './views/home/Home';
import CreateUser from './views/create/CreateUser';
import UpdateUser from './views/update/Update';
import Signup from './views/signup/Signup';
import Login from './views/login/Login';
import Layout from './components/Layout';
import ProtectedRoute from './components/ProtectedRoute';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}>
        {/* public routes */}
        <Route path="signup" element={<Signup />} />
        <Route path="login" element={<Login />} />

        {/* private routes */}
        <Route element={<ProtectedRoute />}>
          <Route path="/" element={<Home />} />
          <Route path="create" element={<CreateUser />} />
          <Route path="update" element={<UpdateUser />} />
        </Route>
      </Route>
    </Routes>
  );
}

export default App;
