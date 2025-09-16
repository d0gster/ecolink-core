const API_BASE = 'http://localhost:8080/api/v1';

export async function createLink(url: string, userID: string) {
	const response = await fetch(`${API_BASE}/links`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'X-User-ID': userID
		},
		body: JSON.stringify({ url })
	});

	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(`Erro ao criar link: ${response.status} ${errorText}`);
	}

	return response.json() as Promise<{ short_url: string; qr_code: string }>;
}

export async function getUserLinks(userID: string) {
	const response = await fetch(`${API_BASE}/links`, {
		headers: {
			'X-User-ID': userID
		}
	});

	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(`Erro ao carregar links: ${response.status} ${errorText}`);
	}

	const data = await (response.json() as Promise<{ links: any[] }>);
	return data.links || [];
}