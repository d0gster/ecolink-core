const API_BASE = 'http://localhost:8080/api/v1';

export async function createLink(url) {
	const response = await fetch(`${API_BASE}/links`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'X-User-ID': 'demo-user'
		},
		body: JSON.stringify({ url })
	});
	
	if (!response.ok) {
		throw new Error('Erro ao criar link');
	}
	
	return response.json();
}

export async function getUserLinks() {
	const response = await fetch(`${API_BASE}/links`, {
		headers: {
			'X-User-ID': 'demo-user'
		}
	});
	
	if (!response.ok) {
		throw new Error('Erro ao carregar links');
	}
	
	const data = await response.json();
	return data.links || [];
}