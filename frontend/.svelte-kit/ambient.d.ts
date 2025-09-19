
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * Environment variables [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env`. Like [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), this module cannot be imported into client-side code. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * _Unlike_ [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), the values exported from this module are statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * ```ts
 * import { API_KEY } from '$env/static/private';
 * ```
 * 
 * Note that all environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * 
 * ```
 * MY_FEATURE_FLAG=""
 * ```
 * 
 * You can override `.env` values from the command line like so:
 * 
 * ```sh
 * MY_FEATURE_FLAG="enabled" npm run dev
 * ```
 */
declare module '$env/static/private' {
	export const VITE_API_URL: string;
	export const VITE_GOOGLE_CLIENT_ID: string;
	export const VITE_GOOGLE_CLIENT_SECRET: string;
	export const AUTH_SECRET: string;
	export const VITE_AUTH0_DOMAIN: string;
	export const VITE_AUTH0_CLIENT_ID: string;
	export const VITE_AUTH0_AUDIENCE: string;
	export const LESSOPEN: string;
	export const VSCODE_CWD: string;
	export const VSCODE_ESM_ENTRYPOINT: string;
	export const LANGUAGE: string;
	export const USER: string;
	export const VSCODE_NLS_CONFIG: string;
	export const npm_config_user_agent: string;
	export const XDG_SEAT: string;
	export const VSCODE_HANDLES_UNCAUGHT_ERRORS: string;
	export const XDG_SESSION_TYPE: string;
	export const npm_node_execpath: string;
	export const SHLVL: string;
	export const LD_LIBRARY_PATH: string;
	export const npm_config_noproxy: string;
	export const HOME: string;
	export const CHROME_DESKTOP: string;
	export const APPDIR: string;
	export const DESKTOP_SESSION: string;
	export const NO_PROXY: string;
	export const npm_package_json: string;
	export const PERLLIB: string;
	export const VSCODE_IPC_HOOK: string;
	export const GTK_MODULES: string;
	export const XDG_SEAT_PATH: string;
	export const LC_MONETARY: string;
	export const npm_config_userconfig: string;
	export const npm_config_local_prefix: string;
	export const DBUS_SESSION_BUS_ADDRESS: string;
	export const GIO_LAUNCHED_DESKTOP_FILE_PID: string;
	export const COLOR: string;
	export const VSCODE_CRASH_REPORTER_PROCESS_TYPE: string;
	export const QT_QPA_PLATFORMTHEME: string;
	export const VSCODE_L10N_BUNDLE_LOCATION: string;
	export const GTK_IM_MODULE: string;
	export const LOGNAME: string;
	export const OWD: string;
	export const _: string;
	export const npm_config_prefix: string;
	export const npm_config_npm_version: string;
	export const XDG_SESSION_CLASS: string;
	export const XDG_SESSION_ID: string;
	export const npm_config_cache: string;
	export const GNOME_DESKTOP_SESSION_ID: string;
	export const npm_config_node_gyp: string;
	export const PATH: string;
	export const SESSION_MANAGER: string;
	export const GDM_LANG: string;
	export const APPIMAGE: string;
	export const GTK3_MODULES: string;
	export const NODE: string;
	export const npm_package_name: string;
	export const XDG_SESSION_PATH: string;
	export const EXPERIMENTAL_HTTP_PROXY_SUPPORT: string;
	export const LC_ADDRESS: string;
	export const VSCODE_PROCESS_TITLE: string;
	export const XDG_RUNTIME_DIR: string;
	export const GDK_BACKEND: string;
	export const DISPLAY: string;
	export const LANG: string;
	export const XDG_CURRENT_DESKTOP: string;
	export const LC_TELEPHONE: string;
	export const XDG_SESSION_DESKTOP: string;
	export const XAUTHORITY: string;
	export const CURSOR_TRACE_ID: string;
	export const npm_lifecycle_script: string;
	export const GSETTINGS_SCHEMA_DIR: string;
	export const XDG_GREETER_DATA_DIR: string;
	export const ORIGINAL_XDG_CURRENT_DESKTOP: string;
	export const SHELL: string;
	export const LC_NAME: string;
	export const ARGV0: string;
	export const npm_package_version: string;
	export const npm_lifecycle_event: string;
	export const QT_ACCESSIBILITY: string;
	export const ELECTRON_RUN_AS_NODE: string;
	export const GDMSESSION: string;
	export const LESSCLOSE: string;
	export const LC_MEASUREMENT: string;
	export const GPG_AGENT_INFO: string;
	export const LC_IDENTIFICATION: string;
	export const XDG_VTNR: string;
	export const npm_config_globalconfig: string;
	export const npm_config_init_module: string;
	export const PWD: string;
	export const npm_execpath: string;
	export const XDG_CONFIG_DIRS: string;
	export const VSCODE_CODE_CACHE_PATH: string;
	export const XDG_DATA_DIRS: string;
	export const npm_config_global_prefix: string;
	export const LC_NUMERIC: string;
	export const npm_command: string;
	export const LC_PAPER: string;
	export const QT_PLUGIN_PATH: string;
	export const VSCODE_PID: string;
	export const INIT_CWD: string;
	export const EDITOR: string;
}

/**
 * Similar to [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private), except that it only includes environment variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Values are replaced statically at build time.
 * 
 * ```ts
 * import { PUBLIC_BASE_URL } from '$env/static/public';
 * ```
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to runtime environment variables, as defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://svelte.dev/docs/kit/cli)), this is equivalent to `process.env`. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * This module cannot be imported into client-side code.
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * console.log(env.DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 * 
 * > [!NOTE] In `dev`, `$env/dynamic` always includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 */
declare module '$env/dynamic/private' {
	export const env: {
		VITE_API_URL: string;
		VITE_GOOGLE_CLIENT_ID: string;
		VITE_GOOGLE_CLIENT_SECRET: string;
		AUTH_SECRET: string;
		VITE_AUTH0_DOMAIN: string;
		VITE_AUTH0_CLIENT_ID: string;
		VITE_AUTH0_AUDIENCE: string;
		LESSOPEN: string;
		VSCODE_CWD: string;
		VSCODE_ESM_ENTRYPOINT: string;
		LANGUAGE: string;
		USER: string;
		VSCODE_NLS_CONFIG: string;
		npm_config_user_agent: string;
		XDG_SEAT: string;
		VSCODE_HANDLES_UNCAUGHT_ERRORS: string;
		XDG_SESSION_TYPE: string;
		npm_node_execpath: string;
		SHLVL: string;
		LD_LIBRARY_PATH: string;
		npm_config_noproxy: string;
		HOME: string;
		CHROME_DESKTOP: string;
		APPDIR: string;
		DESKTOP_SESSION: string;
		NO_PROXY: string;
		npm_package_json: string;
		PERLLIB: string;
		VSCODE_IPC_HOOK: string;
		GTK_MODULES: string;
		XDG_SEAT_PATH: string;
		LC_MONETARY: string;
		npm_config_userconfig: string;
		npm_config_local_prefix: string;
		DBUS_SESSION_BUS_ADDRESS: string;
		GIO_LAUNCHED_DESKTOP_FILE_PID: string;
		COLOR: string;
		VSCODE_CRASH_REPORTER_PROCESS_TYPE: string;
		QT_QPA_PLATFORMTHEME: string;
		VSCODE_L10N_BUNDLE_LOCATION: string;
		GTK_IM_MODULE: string;
		LOGNAME: string;
		OWD: string;
		_: string;
		npm_config_prefix: string;
		npm_config_npm_version: string;
		XDG_SESSION_CLASS: string;
		XDG_SESSION_ID: string;
		npm_config_cache: string;
		GNOME_DESKTOP_SESSION_ID: string;
		npm_config_node_gyp: string;
		PATH: string;
		SESSION_MANAGER: string;
		GDM_LANG: string;
		APPIMAGE: string;
		GTK3_MODULES: string;
		NODE: string;
		npm_package_name: string;
		XDG_SESSION_PATH: string;
		EXPERIMENTAL_HTTP_PROXY_SUPPORT: string;
		LC_ADDRESS: string;
		VSCODE_PROCESS_TITLE: string;
		XDG_RUNTIME_DIR: string;
		GDK_BACKEND: string;
		DISPLAY: string;
		LANG: string;
		XDG_CURRENT_DESKTOP: string;
		LC_TELEPHONE: string;
		XDG_SESSION_DESKTOP: string;
		XAUTHORITY: string;
		CURSOR_TRACE_ID: string;
		npm_lifecycle_script: string;
		GSETTINGS_SCHEMA_DIR: string;
		XDG_GREETER_DATA_DIR: string;
		ORIGINAL_XDG_CURRENT_DESKTOP: string;
		SHELL: string;
		LC_NAME: string;
		ARGV0: string;
		npm_package_version: string;
		npm_lifecycle_event: string;
		QT_ACCESSIBILITY: string;
		ELECTRON_RUN_AS_NODE: string;
		GDMSESSION: string;
		LESSCLOSE: string;
		LC_MEASUREMENT: string;
		GPG_AGENT_INFO: string;
		LC_IDENTIFICATION: string;
		XDG_VTNR: string;
		npm_config_globalconfig: string;
		npm_config_init_module: string;
		PWD: string;
		npm_execpath: string;
		XDG_CONFIG_DIRS: string;
		VSCODE_CODE_CACHE_PATH: string;
		XDG_DATA_DIRS: string;
		npm_config_global_prefix: string;
		LC_NUMERIC: string;
		npm_command: string;
		LC_PAPER: string;
		QT_PLUGIN_PATH: string;
		VSCODE_PID: string;
		INIT_CWD: string;
		EDITOR: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * Similar to [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), but only includes variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Note that public dynamic environment variables must all be sent from the server to the client, causing larger network requests — when possible, use `$env/static/public` instead.
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.PUBLIC_DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
