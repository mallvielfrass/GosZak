define(['uikit'], function(UIkit) {
	class SearchCard extends HTMLElement {
		constructor() {
			super();

			if(this.item === null) {
				return;
			}

			let shadowRoot = this.attachShadow({mode: 'open'});

			let actionsHTML = '<ul class="uk-list">';

			for(let action of this.item.Actions) {
				actionsHTML += `<li><a href="${action.Link}" target="_blank" class="uk-link-text">${action.Name}</a></li>`;
			}

			actionsHTML += '</ul>';

			shadowRoot.innerHTML = `
				<link rel="stylesheet" href="/css/uikit.min.css" />
				<style>
					.search-card-title, .search-card-body {
						word-wrap: break-word;
					}
				</style>
				<div class="uk-card uk-card-default">
					<div class="uk-card-header">
						<h3 class="uk-card-title uk-margin-remove-bottom search-card-title"><a href="${this.item.Link}" target="_blank" class="uk-link-heading">${this.item.Name}</a></h3>
						<p class="uk-text-meta uk-margin-remove-top">
							Размещено: ${new Date(this.item.PublishDate * 1000).toLocaleDateString()}
							<br>
							Обновлено: ${new Date(this.item.UpdateDate * 1000).toLocaleDateString()}
						</p>
					</div>
					<div class="search-card-body uk-card-body">
						<p>
							<b>Заказчик:</b> <a href="${this.item.CustomerLink}" target="_blank">${this.item.Customer}</a>
						</p>
						<p>
							<b>Тип:</b> ${this.item.Type}
							<br>
							<b>Статус:</b> ${this.item.Status} / ${this.item.Law}
							${this.item.Price.length > 0 ? `
								<br>
								<b>Начальная цена:</b> ${this.item.Price} (${this.item.Currency})`
							: ''}
						</p>
						<p>
						${this.item.Description}
						</p>
						${this.item.Lots.length > 0 ?
							`<div class="uk-margin-top uk-text-center">
								<button type="button" id="search-lot-button" class="uk-button uk-button-default">Информация о лотах</button>
							</div>
							<div id="search-lot-content" class="uk-margin-top" hidden>
								<ul class="uk-list uk-list-divider uk-list-large">
									<li>
										${this.item.Lots.map(lotItem => 
											`<p>
												<b>${lotItem.Name}</b> ${lotItem.Description}
											</p>
											<p>
												<i>Начальная цена: <b>${lotItem.Price}</b> (${lotItem.Currency})</i>
											</p>`
										).join('</li><li>')}
									</li>
								</ul>
							</div>`
						: ''}
						${this.item.Ids.length > 0 ?
							`<div class="uk-margin-top uk-text-center">
								<button type="button" id="search-id-button" class="uk-button uk-button-default">Идентификационные коды закупки</button>
							</div>
							<div id="search-id-content" class="uk-margin-top" hidden>${this.item.Ids.join('<br>')}</div>`
						: ''}
					</div>
					<div class="uk-card-footer">
						${actionsHTML}
					</div>
				</div>
			`;

			let lotButton = shadowRoot.querySelector('#search-lot-button');
			let idButton = shadowRoot.querySelector('#search-id-button');

			let lotToggle = UIkit.toggle(shadowRoot.querySelector('#search-lot-content'));
			let idToggle = UIkit.toggle(shadowRoot.querySelector('#search-id-content'));

			if(lotButton !== null) {
				lotButton.addEventListener('click', () => {
					lotToggle.toggle();
				});
			}

			if(idButton !== null) {
				idButton.addEventListener('click', () => {
					idToggle.toggle();
				});
			}
		}

		get item() {
			return JSON.parse(this.getAttribute('item') || 'null');
		}
	}

	customElements.define('search-card', SearchCard);
});
