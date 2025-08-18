import { API } from './services/API.js';
import { HomePage } from './components/HomePage.js';
import { MovieDetailsPage } from './components/MovieDetailsPage.js';
import './components/AnimatedLoading.js';
import './components/YouTubeEmbed.js';
import { Router } from './services/Router.js';

window.addEventListener('DOMContentLoaded', (event) => {
  app.Router.init();
});

window.app = {
  search: (event) => {
    event.preventDefault();

    const keywords = document.querySelector('input[type=search]').value;

    // TODO
    console.log('Keywords', keywords);
  },
  api: API,
  Router,
};
