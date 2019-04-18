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

		library['searchQuery'] = searchForm.elements['searchQuery'].value;

		let response = await fetch('/api/search', {
			method: 'POST',
			body: JSON.stringify(library)
		});
	
		document.querySelector('#result').innerHTML = await response.text();
	}

	if(location.pathname === '/' || location.pathname === '/index.html') {
		let searchForm = document.querySelector('#searchForm');

		searchForm.onsubmit = () => {
			sendSearchForm();
			return false;
		};
	}
});