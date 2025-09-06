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
    const q = document.querySelector('input[type=search]').value;
    app.Router.go('/movies?q=' + q);
  },
  api: API,
  Router,
};
