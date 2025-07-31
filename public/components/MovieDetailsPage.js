import { API } from '../services/API.js';
import { MovieItem } from './MovieItem.js';

export class MovieDetailsPage extends HTMLElement {
  id = null;
  movie = null;

  async render() {
    try {
      this.movie = await API.getMovieById(this.id);
    } catch (err) {
      // TODO: alert user
      console.error(err);
    }

    const template = document.getElementById('template-movie-details');
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    this.querySelector('h2').textContent = this.movie.title;
    this.querySelector('h3').textContent = this.movie.tagline;
    this.querySelector('h3').textContent = this.movie.tagline;
    this.querySelector('img').src = this.movie.poster_url;
  }

  connectedCallback() {
    this.id = 14;
    this.render();
  }
}

customElements.define('movie-details-page', MovieDetailsPage);
