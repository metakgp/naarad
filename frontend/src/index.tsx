/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';

const root = document.getElementById('root');

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    'Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?',
  );
}

const redirect = sessionStorage.getItem("naarad-spa-redirect");

if (redirect) {
  sessionStorage.removeItem("naarad-spa-redirect");
  window.history.replaceState(null, "", redirect);
}

render(() => <App />, root!);
