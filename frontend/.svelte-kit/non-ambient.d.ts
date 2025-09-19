
// this file is generated — do not edit it


declare module "svelte/elements" {
	export interface HTMLAttributes<T> {
		'data-sveltekit-keepfocus'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-noscroll'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-preload-code'?:
			| true
			| ''
			| 'eager'
			| 'viewport'
			| 'hover'
			| 'tap'
			| 'off'
			| undefined
			| null;
		'data-sveltekit-preload-data'?: true | '' | 'hover' | 'tap' | 'off' | undefined | null;
		'data-sveltekit-reload'?: true | '' | 'off' | undefined | null;
		'data-sveltekit-replacestate'?: true | '' | 'off' | undefined | null;
	}
}

export {};


declare module "$app/types" {
	export interface AppTypes {
		RouteId(): "/" | "/auth" | "/auth/callback" | "/auth/callback/google" | "/dashboard" | "/profile" | "/result";
		RouteParams(): {
			
		};
		LayoutParams(): {
			"/": Record<string, never>;
			"/auth": Record<string, never>;
			"/auth/callback": Record<string, never>;
			"/auth/callback/google": Record<string, never>;
			"/dashboard": Record<string, never>;
			"/profile": Record<string, never>;
			"/result": Record<string, never>
		};
		Pathname(): "/" | "/auth" | "/auth/" | "/auth/callback" | "/auth/callback/" | "/auth/callback/google" | "/auth/callback/google/" | "/dashboard" | "/dashboard/" | "/profile" | "/profile/" | "/result" | "/result/";
		ResolvedPathname(): `${"" | `/${string}`}${ReturnType<AppTypes['Pathname']>}`;
		Asset(): "/images/download.png" | "/videos/README.md" | "/videos/eco-background-desktop.mp4" | string & {};
	}
}