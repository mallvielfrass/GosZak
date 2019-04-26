define(['uikit', 'search'], function(UIkit, Search) {
	class SearchPager extends HTMLElement {
		constructor() {
			super();

			if(isNaN(this.page) || isNaN(this.totalPage)) {
				return;
			}

			let shadowRoot = this.attachShadow({mode: 'open'});

			let minPage = this.page;
			let maxPage = this.page;
	
			if(this.page > 1) {
				minPage = this.page-1;
			}
	
			if(this.page > 2) {
				minPage = this.page-2;
			}

			shadowRoot.innerHTML = `
				<link rel="stylesheet" href="/css/uikit.min.css" />
				${this.totalPage > 1 ?
					`<ul class="uk-pagination uk-flex-center" uk-margin>
						${this.page > 1 ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page-1}"><span id="previous-page" uk-pagination-previous></span></a></li>`
						: ''}
						${minPage != 1 ?
							`<li><a href="#" class="pager" data-pagenumber="1">1</a></li>
							${minPage > 2 ?
								`<li class="uk-disabled"><span>...</span></li>`
							: ''}`
						: ''}
						${this.page > 2 ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page-2}">${this.page-2}</a></li>`
						: ''}
						${this.page > 1 ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page-1}">${this.page-1}</a></li>`
						: ''}
						<li class="uk-active"><span>${this.page}</span></li>
						${this.page+1 <= this.totalPage ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page+1}">${maxPage = this.page+1}</a></li>`
						: ''}
						${this.page+2 <= this.totalPage ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page+2}">${maxPage = this.page+2}</a></li>`
						: ''}
						${this.page < 3 && this.page+3 < this.totalPage ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page+3}">${maxPage = this.page+3}</a></li>`
						: ''}
						${this.page < 2 && this.page+4 < this.totalPage ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page+4}">${maxPage = this.page+4}</a></li>`
						: ''}
						${maxPage != this.totalPage ?
							`${maxPage < this.totalPage-1 ?
								`<li class="uk-disabled"><span>...</span></li>`
							: ''}
								<li><a href="#" class="pager" data-pagenumber="${this.totalPage}">${this.totalPage}</a></li>`
						: ''}
						${this.page < this.totalPage ?
							`<li><a href="#" class="pager" data-pagenumber="${this.page+1}"><span id="next-page" uk-pagination-next></span></a></li>`
						: ''}
					</ul>`
				: ''}
			`;

			let previousPageIcon = UIkit.icon(shadowRoot.querySelector('#previous-page'));

			if(typeof previousPageIcon != 'undefined') {
				previousPageIcon.icon = 'pagination-previous';
				previousPageIcon.$el.classList.add('uk-pagination-previous');
				previousPageIcon.$el.classList.add('uk-icon');

				previousPageIcon.getSvg().then(svg => {
					previousPageIcon.$el.appendChild(svg);
				});
			}

			let nextPageIcon = UIkit.icon(shadowRoot.querySelector('#next-page'));

			if(typeof nextPageIcon != 'undefined') {
				nextPageIcon.icon = 'pagination-next';
				nextPageIcon.$el.classList.add('uk-pagination-next');
				nextPageIcon.$el.classList.add('uk-icon');

				nextPageIcon.getSvg().then(svg => {
					nextPageIcon.$el.appendChild(svg);
				});
			}

			for(let item of shadowRoot.querySelectorAll('.pager')) {
				item.addEventListener('click', (e) => {
					e.preventDefault();
					document.querySelector('#searchForm').elements['pageNumber'].value = item.getAttribute('data-pagenumber');
					Search.start();
				});
			}
		}

		get page() {
			return Number(this.getAttribute('page') || NaN);
		}

		get totalPage() {
			return Number(this.getAttribute('total-page') || NaN);
		}
	}

	customElements.define('search-pager', SearchPager);
});
