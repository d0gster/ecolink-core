

export const index = 6;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/dashboard/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/6.90db0908.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/index.02aeeab6.js","_app/immutable/chunks/each.e59479a4.js","_app/immutable/chunks/auth-guard.a0cccbed.js","_app/immutable/chunks/navigation.f767a769.js","_app/immutable/chunks/singletons.edde51cd.js","_app/immutable/chunks/env.0a8cfa1c.js"];
export const stylesheets = [];
export const fonts = [];
