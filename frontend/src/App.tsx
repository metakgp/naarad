import type { Component } from 'solid-js'
import { Route, Router } from '@solidjs/router';
import { Register } from './pages/Register';
import './styles/index.scss';

const App: Component = () => {
  return (
    <div>
      <Router>
        <Route path="/" component={Register} />
      </Router>
    </div>
  );
};

export default App;
