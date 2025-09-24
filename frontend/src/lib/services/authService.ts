import { browser } from '$app/environment';
import { config } from '$lib/config/env';
import { isLoading, isAuthenticated, user as userStore } from '$lib/stores/auth';

import type { User } from '$lib/stores/auth';

/**
 * Responsibility: encapsulate all authentication-related logic for the frontend.
 * - verifySession(): revalidates session using backend /api/v1/me
 * - handleCallback(): exchange code via backend (delegates to existing endpoint)
 * - login(), logout()
 *
 * This module follows SRP and keeps UI components free of business logic.
 */
export async function verifySession(): Promise<User | null> {
  if (!browser) return null;

  isLoading.set(true);

  try {
    const response = await fetch(`${config.apiUrl}/api/v1/me`, {
      credentials: 'include'
    });

    if (!response.ok) {
      userStore.set(null);
      isAuthenticated.set(false);
      return null;
    }

    const data: User = await response.json();
    userStore.set(data);
    isAuthenticated.set(true);
    return data;
  } catch (error) {
    console.error('verifySession error:', error);
    userStore.set(null);
    isAuthenticated.set(false);
    return null;
  } finally {
    isLoading.set(false);
  }
}

export async function handleCallback(code: string, state?: string) {
  // Generate state if not provided (for security)
  if (!state) {
    state = generateSecureState();
  }

  // Delegate to backend auth callback endpoint
  const response = await fetch(`${config.apiUrl}/auth/google/callback`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ 
      code, 
      redirect_uri: `${window.location.origin}/auth/callback/google`, 
      state 
    })
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.error || 'Authentication failed');
  }

  const data = await response.json();
  userStore.set(data.user);
  isAuthenticated.set(true);

  return data;
}

export function login(state?: string) {
  if (!browser) return;

  // Always generate a secure state parameter for CSRF protection
  if (!state) {
    state = generateSecureState();
  }

  // Store state in sessionStorage for verification
  sessionStorage.setItem('oauth_state', state);

  const params = new URLSearchParams({
    client_id: config.googleClientId,
    redirect_uri: `${window.location.origin}/auth/callback/google`,
    response_type: 'code',
    scope: 'openid profile email',
    state: state
  });

  window.location.href = `https://accounts.google.com/o/oauth2/v2/auth?${params}`;
}

export async function logout() {
  await fetch(`${config.apiUrl}/auth/logout`, {
    method: 'POST',
    credentials: 'include'
  }).catch(() => {});

  userStore.set(null);
  isAuthenticated.set(false);
  if (browser) window.location.href = '/';
}

// Generate cryptographically secure state parameter
function generateSecureState(): string {
  const array = new Uint8Array(32);
  crypto.getRandomValues(array);
  return btoa(String.fromCharCode(...array))
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
}

export default {
  verifySession,
  handleCallback,
  login,
  logout
};


