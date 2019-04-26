define(function() {
	class AppHeader extends HTMLElement {
		constructor() {
			super();

			let shadowRoot = this.attachShadow({mode: 'open'});

			shadowRoot.innerHTML = `
				<link rel="stylesheet" href="/css/uikit.min.css" />
				<nav class="uk-navbar-container" uk-navbar>
					<div class="uk-navbar-left">
						<a class="uk-navbar-item uk-logo" href="/">Парсер госзакупок</a>
					</div>
				</nav>
			`;
		}
	}

	customElements.define('app-header', AppHeader);
});
