requirejs(['uikit', 'search', 'app-header', 'app-footer', 'search-card', 'search-pager'], function(UIkit, Search) {
	if(location.pathname === '/' || location.pathname === '/index.html') {
		let searchForm = document.querySelector('#searchForm');

		searchForm.addEventListener('submit', (e) => {
			e.preventDefault();
			searchForm.elements['pageNumber'].value = 1;
			Search.start();
		});
	}
});
