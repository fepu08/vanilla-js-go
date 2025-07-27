import { API } from '../services/API.js';

export class HomePage extends HTMLElement {
  async render() {
    const topMovies = await API.getTopMovies();
    renderMovies(topMovies, document.querySelector('#top-10 ul'));

    const randomMovies = await API.randomMovies();
    renderMovies(randomMovies, document.querySelector('#random ul'));

    function renderMovies(movies, ul) {
      ul.innerHTML = '';
      movies.forEach((m) => {
        const li = document.createElement('li');
        li.innerHTML = m.title;
        ul.appendChild(li);
      });
    }
  }

  connectedCallback() {
    const template = document.getElementById('template-home');
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    this.render();
  }
}

customElements.define('home-page', HomePage);
