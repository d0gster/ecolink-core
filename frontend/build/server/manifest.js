const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["images/download.png","videos/README.md","videos/eco-background-desktop.mp4"]),
	mimeTypes: {".png":"image/png",".md":"text/markdown",".mp4":"video/mp4"},
	_: {
		client: {"start":"_app/immutable/entry/start.d8476eb4.js","app":"_app/immutable/entry/app.630e0e6d.js","imports":["_app/immutable/entry/start.d8476eb4.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/singletons.edde51cd.js","_app/immutable/entry/app.630e0e6d.js","_app/immutable/chunks/preload-helper.a4192956.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/index.02aeeab6.js"],"stylesheets":[],"fonts":[]},
		nodes: [
			__memo(() => import('./chunks/0-c36f2e09.js')),
			__memo(() => import('./chunks/1-3d0d60ca.js')),
			__memo(() => import('./chunks/2-f94c4c5e.js')),
			__memo(() => import('./chunks/3-11e45226.js')),
			__memo(() => import('./chunks/4-57b3fd47.js')),
			__memo(() => import('./chunks/5-785ceff4.js')),
			__memo(() => import('./chunks/6-3ae64bfe.js')),
			__memo(() => import('./chunks/7-f349b711.js')),
			__memo(() => import('./chunks/8-15cfd354.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/auth",
				pattern: /^\/auth\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/auth/callback",
				pattern: /^\/auth\/callback\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/auth/callback/google",
				pattern: /^\/auth\/callback\/google\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/dashboard",
				pattern: /^\/dashboard\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/profile",
				pattern: /^\/profile\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/result",
				pattern: /^\/result\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 8 },
				endpoint: null
			}
		],
		matchers: async () => {
			
			return {  };
		}
	}
}
})();

const prerendered = new Set([]);

export { manifest, prerendered };
//# sourceMappingURL=manifest.js.map
