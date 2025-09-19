import { config } from '$lib/config/env';

export async function createLink(url: string) {
	const response = await fetch(`${config.apiUrl}/api/v1/links`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		credentials: 'include',
		body: JSON.stringify({ url })
	});

	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(`Error creating link: ${response.status} ${errorText}`);
	}

	return response.json();
}

export async function getUserLinks() {
	const response = await fetch(`${config.apiUrl}/api/v1/links`, {
		credentials: 'include'
	});

	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(`Error loading links: ${response.status} ${errorText}`);
	}

	const data = await response.json();
	return data.links || [];
}