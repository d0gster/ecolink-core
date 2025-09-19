

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.fb1c9483.js","_app/immutable/chunks/preload-helper.a4192956.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/index.02aeeab6.js","_app/immutable/chunks/navigation.f767a769.js","_app/immutable/chunks/singletons.edde51cd.js","_app/immutable/chunks/env.0a8cfa1c.js"];
export const stylesheets = [];
export const fonts = [];
