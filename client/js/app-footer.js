define(function() {
	class AppFooter extends HTMLElement {
		constructor() {
			super();

			let shadowRoot = this.attachShadow({mode: 'open'});

			shadowRoot.innerHTML = `
				<link rel="stylesheet" href="/css/uikit.min.css" />
				<div class="uk-section-small">
					<div class="uk-container uk-container-expand uk-text-center uk-position-relative">
						<ul class="uk-subnav uk-flex-inline uk-flex-center uk-margin-remove-bottom" uk-margin>
							<li>
								<a href="/about.html">О парсере</a>
							</li>
						</ul>
					</div>
				</div>
			`;
		}
	}

	customElements.define('app-footer', AppFooter);
});
