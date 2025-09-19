import { c as create_ssr_component, a as subscribe } from './ssr-98fcd782.js';
import { p as pendingLink, u as user, a as isAuthenticated } from './auth-aa1f0a46.js';
import { p as page } from './stores-7d44a32b.js';
import './index-a2bcafe3.js';

const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$unsubscribe_page;
  let $$unsubscribe_pendingLink;
  let $$unsubscribe_user;
  let $$unsubscribe_isAuthenticated;
  $$unsubscribe_page = subscribe(page, (value) => value);
  $$unsubscribe_pendingLink = subscribe(pendingLink, (value) => value);
  $$unsubscribe_user = subscribe(user, (value) => value);
  $$unsubscribe_isAuthenticated = subscribe(isAuthenticated, (value) => value);
  $$unsubscribe_page();
  $$unsubscribe_pendingLink();
  $$unsubscribe_user();
  $$unsubscribe_isAuthenticated();
  return `${$$result.head += `<!-- HEAD_svelte-1j3k86_START -->${$$result.title = `<title>Link Created - EcoLink</title>`, ""}<!-- HEAD_svelte-1j3k86_END -->`, ""} <div class="text-center"><div class="max-w-3xl mx-auto">${`<div class="py-12" data-svelte-h="svelte-xh5kmu"><div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-eco-600"></div> <p class="mt-4 text-gray-600">Creating your sustainable link...</p></div>`}</div></div>`;
});

export { Page as default };
//# sourceMappingURL=_page.svelte-18a34fd2.js.map
