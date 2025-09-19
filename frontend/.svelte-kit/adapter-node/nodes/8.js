

export const index = 8;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/result/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/8.bcddccc3.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/index.02aeeab6.js","_app/immutable/chunks/navigation.f767a769.js","_app/immutable/chunks/singletons.edde51cd.js","_app/immutable/chunks/stores.f244effd.js"];
export const stylesheets = [];
export const fonts = [];
