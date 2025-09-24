export const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

// Constantes para validação de URLs
export const URL_PATTERNS = {
	HTTP: /^https?:\/\//i,
	VALID_DOMAIN: /^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]?\.[a-zA-Z]{2,}$/
} as const;

// Constantes para validação de entrada
export const VALIDATION = {
	MAX_URL_LENGTH: 2048,
	MIN_URL_LENGTH: 10
} as const;

