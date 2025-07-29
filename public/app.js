import { API } from './services/API.js';
import { HomePage } from './components/HomePage.js';
import './components/AnimatedLoading.js';

window.addEventListener('DOMContentLoaded', (event) => {
  document.querySelector('main').appendChild(new HomePage());
});

window.app = {
  search: (event) => {
    event.preventDefault();

    const keywords = document.querySelector('input[type=search]').value;

    // TODO
    console.log('Keywords', keywords);
  },
  api: API,
};
