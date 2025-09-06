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
  showError: (message = 'There was an error.', goToHome = true) => {
    document.getElementById('alert-modal').showModal();
    document.querySelector('#alert-modal p').textContent = message;
    if (goToHome) app.Router.go('/');
  },
  closeError: () => {
    document.getElementById('alert-modal').close();
  },
  search: (event) => {
    event.preventDefault();
    const q = document.querySelector('input[type=search]').value;
    if (q) {
      app.Router.go('/movies?q=' + q);
    }
  },
  searchOrderChange: (order) => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get('q');
    const genre = urlParams.get('genre') ?? '';
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  searchFilterChange: (genre) => {
    const urlParams = new URLSearchParams(window.location.search);
    const q = urlParams.get('q');
    const order = urlParams.get('order') ?? '';
    app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
  },
  api: API,
  Router,
};
