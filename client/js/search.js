define(function() {
	class Search {
		async sendForm() {
			let library = {};
			let searchForm = document.forms['searchForm']; 
			let lawNumber = searchForm.elements['lawNumber'];
			let procedureStatus = searchForm.elements['procedureStatus'];

			for(let i = 0; i < lawNumber.length; i++) {
				if(lawNumber[i].checked === true) {
					if(typeof library['lawNumber'] === 'undefined') {
						library['lawNumber'] = [];
					}

					library['lawNumber'].push(lawNumber[i].value);
				}
			}

			for(let i = 0; i < procedureStatus.length; i++) {
				if(procedureStatus[i].checked === true) {
					if(typeof library['procedureStatus'] === 'undefined') {
						library['procedureStatus'] = [];
					}

					library['procedureStatus'].push(procedureStatus[i].value);
				}
			}

			if(searchForm.elements['sortDirection'].value !== '') {
				library['sortDirection'] = searchForm.elements['sortDirection'].value;
			}

			if(searchForm.elements['sortBy'].value !== '') {
				library['sortBy']= searchForm.elements['sortBy'].value;
			}

			if(searchForm.elements['cityName'].value !== '') {
				library['cityName'] = searchForm.elements['cityName'].value;
			}

			if(searchForm.elements['publishDateFrom'].value !== '') {
				library['publishDateFrom'] = Math.round(Date.parse(searchForm.elements['publishDateFrom'].value) / 1000);
			}

			if(searchForm.elements['publishDateTo'].value !== '') {
				library['publishDateTo'] = Math.round(Date.parse(searchForm.elements['publishDateTo'].value) / 1000);
			}

			if(searchForm.elements['pageNumber'].value !== '') {
				library['pageNumber'] = Number(searchForm.elements['pageNumber'].value);
			}

			library['searchString'] = searchForm.elements['searchString'].value;

			let response = await fetch('/api/search', {
				method: 'POST',
				body: JSON.stringify(library)
			});

			return response.json();
		}

		parseResult(result) {
			if(typeof result.Error !== 'undefined') {
				return `<div class="uk-text-center">${result.Error
					.replace(/&/g, '&amp;')
					.replace(/</g, '&lt;')
					.replace(/>/g, '&gt;')
					.replace(/"/g, '&quot;')}</div>`
			}

			let itemsHTML = '';

			for(let item of result.Items) {
				let searchCard = document.createElement('search-card');
				searchCard.setAttribute('item', JSON.stringify(item));

				itemsHTML += searchCard.outerHTML;
			}

			let returnHTML = `
				<div class="uk-child-width-1-2@s uk-grid-match" uk-grid>
					${itemsHTML}
				</div>
				<search-pager page="${result.Page}" total-page="${result.TotalPage}"></search-pager>
				<div class="uk-margin uk-text-center">Всего записей: <b>${result.Total}</b></div>
			`;

			return returnHTML;
		}

		async start() {
			document.querySelector('#result').innerHTML = this.parseResult(await this.sendForm());
			location.href = '#result';
		}
	}

	return new Search;
});
