import authService from '$lib/services/authService';

export const initGoogleAuth = authService.verifySession;
export const login = authService.login;
export const handleCallback = authService.handleCallback;
export const logout = authService.logout;