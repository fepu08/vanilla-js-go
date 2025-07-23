import { API } from './services/API.js';

window.app = {
  search: (event) => {
    event.preventDefault();

    const keywords = document.querySelector('input[type=search]').value;

    // TODO
    console.log('Keywords', keywords);
  },
  api: API,
};
