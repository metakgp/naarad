import type { Component } from 'solid-js'
import logo from './logo.svg';
import styles from './App.module.css';
import { Route, Router } from '@solidjs/router';
import { Register } from './pages/Register';

const App: Component = () => {
  return (
    <div class={styles.App}>
      <Router>
        <Route path="/" component={Register} />
      </Router>
    </div>
  );
};

export default App;
