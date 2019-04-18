requirejs(['uikit'], function() {
	async function sendSearchForm() {
		let library = {};
		let searchForm = document.forms['searchForm']; 
		let lawNumber = searchForm.elements['lawNumber'];
		let procedureType = searchForm.elements['procedureType'];

		document.querySelector('#result').innerText = '';

		for(let i = 0; i < lawNumber.length; i++) {
			if(lawNumber[i].checked === true) {
				if(typeof library['lawNumber'] === 'undefined') {
					library['lawNumber'] = [];
				}

				library['lawNumber'].push(lawNumber[i].value);
			}
		}

		for(let i = 0; i < procedureType.length; i++) {
			if(procedureType[i].checked === true) {
				if(typeof library['procedureType'] === 'undefined') {
					library['procedureType'] = [];
				}

				library['procedureType'].push(procedureType[i].value);
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

		library['searchString'] = searchForm.elements['searchString'].value;

		let response = await fetch('/api/search', {
			method: 'POST',
			body: JSON.stringify(library)
		});
	
		document.querySelector('#result').innerHTML = parseSearchResult(await response.text());
	}

	function parseSearchResult(result) {
		try {
			let resultJSON = JSON.parse(result);

			if(typeof resultJSON.Error !== 'undefined') {
				return `<div class="uk-text-center">${resultJSON.Error}</div>`
			}
		} catch(e) {}

		let resultHTML = new DOMParser().parseFromString(result, 'text/html');
		let boxIcons = resultHTML.querySelectorAll('body > .boxIcons');
		let resultTables = resultHTML.querySelectorAll('body > table');
		let allRecords = resultHTML.querySelector('.allRecords');

		allRecords.className += ' uk-width-1-1 uk-text-center';

		for(let i = 0; i < boxIcons.length; i++) {
			boxIcons[i].parentNode.removeChild(boxIcons[i]);
		}

		for(let i = 0; i < resultTables.length; i++) {
			let itemHeader = resultTables[i].querySelector('.descriptTenderTd > dl > dt');

			let itemId = itemHeader.querySelector('a');
			itemId.className = 'uk-link-heading';

			let itemIdUrl = new URL(itemId.href);
			itemIdUrl.origin = 'http://zakupki.gov.ru';
			itemId.href = itemIdUrl.href;

			let itemIdHTML = itemId.outerHTML;
			itemId.parentNode.removeChild(itemId);

			let itemOrganization = resultTables[i].querySelector('.descriptTenderTd > dl > .nameOrganization');
			let organizationLink = itemOrganization.querySelector('a');
			organizationLink.className = 'uk-link-text';

			let organizationUrl = new URL(organizationLink.href);
			organizationUrl.origin = 'http://zakupki.gov.ru';
			organizationLink.href = organizationUrl.href;
			organizationLink.removeAttribute('onclick');

			let organizationLinkHTML = organizationLink.outerHTML;
			itemOrganization.parentNode.removeChild(itemOrganization);

			let itemBody = resultTables[i].querySelectorAll('.descriptTenderTd > dl > *');
			let itemBodyHTML = '';

			for(let i = 0; i < itemBody.length; i++) {
				if(i > 0) {
					itemBodyHTML += '<br>';
				}

				itemBodyHTML += itemBody[i].innerHTML;
			}

			let itemInfo = resultTables[i].querySelectorAll('.tenderTd > dl > *');

			let amountTd = resultTables[i].querySelector('.amountTenderTd');
			amountTd.querySelector('ul').className = 'uk-list';

			let reportBox = resultTables[i].nextElementSibling;
			let reportBoxList = reportBox.querySelector('ul > ul');
			reportBoxList.className = 'uk-list';

			let reportBoxLinks = reportBoxList.querySelectorAll('a');

			for(let i = 0; i < reportBoxLinks.length; i++) {
				reportBoxLinks[i].className = 'uk-link-text';

				if(reportBoxLinks[i].href.length === 0) {
					reportBoxLinks[i].parentNode.removeChild(reportBoxLinks[i]);
					continue;
				}

				let reportBoxUrl = new URL(reportBoxLinks[i].href);
				reportBoxUrl.origin = 'http://zakupki.gov.ru';
				reportBoxLinks[i].href = reportBoxUrl.href;
			}

			amountTd.insertBefore(reportBoxList, amountTd.firstChild);
			reportBox.parentNode.removeChild(reportBox);

			resultTables[i].outerHTML = `<div class="uk-card uk-card-default uk-card-body uk-width-1-2@m">
				<div class="uk-card-header">
					<h3 class="uk-card-title uk-margin-remove-bottom">${itemIdHTML}</h3>
					<p class="uk-text-meta uk-margin-remove-top">
						${itemInfo[0].innerText}
						<br>
						${itemInfo[2].innerText}
						${itemInfo[3].querySelector('strong') ? `<br>
							Начальная цена: ${itemInfo[3].querySelector('strong').innerText} (${itemInfo[3].querySelector('.currency').innerText})`
						: ''}
						<br>
						Заказчик: ${organizationLinkHTML}
					</p>
				</div>
				<div class="uk-card-body">
					${itemBodyHTML}
				</div>
				<div class="uk-card-footer">
					${amountTd !== null ? amountTd.innerHTML : ''}
				</div>
			</div>`;
		}

		returnHTML = `<div class="uk-grid-small" uk-grid>${resultHTML.body.innerHTML}</div>`;

		return returnHTML;
	}

	if(location.pathname === '/' || location.pathname === '/index.html') {
		let searchForm = document.querySelector('#searchForm');

		searchForm.onsubmit = () => {
			sendSearchForm();
			return false;
		};
	}
});