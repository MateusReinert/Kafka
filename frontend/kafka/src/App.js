import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import ProducerPage from './shared/ProducerPage';
import MessagePage from './shared/MessagePage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<MessagePage />} />
        <Route path="/producer1" element={<ProducerPage />} />
      </Routes>
    </Router>
  );
}

export default App;
