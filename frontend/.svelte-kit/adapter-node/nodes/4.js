

export const index = 4;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/auth/callback/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/4.5601018a.js","_app/immutable/chunks/scheduler.ffbd56c5.js","_app/immutable/chunks/index.02aeeab6.js","_app/immutable/chunks/navigation.f767a769.js","_app/immutable/chunks/singletons.edde51cd.js","_app/immutable/chunks/stores.f244effd.js","_app/immutable/chunks/google-direct.658b226e.js","_app/immutable/chunks/authService.f7fb7133.js","_app/immutable/chunks/env.0a8cfa1c.js"];
export const stylesheets = [];
export const fonts = [];
