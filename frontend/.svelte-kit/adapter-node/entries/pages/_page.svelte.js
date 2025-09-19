import { c as create_ssr_component, b as subscribe, a as add_attribute, e as escape } from "../../chunks/ssr.js";
import { u as user, i as isAuthenticated } from "../../chunks/auth.js";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $$unsubscribe_user;
  let $$unsubscribe_isAuthenticated;
  $$unsubscribe_user = subscribe(user, (value) => value);
  $$unsubscribe_isAuthenticated = subscribe(isAuthenticated, (value) => value);
  let url = "";
  $$unsubscribe_user();
  $$unsubscribe_isAuthenticated();
  return `${$$result.head += `<!-- HEAD_svelte-ekfoov_START -->${$$result.title = `<title>EcoLink - Sustainable Shortener</title>`, ""}<!-- HEAD_svelte-ekfoov_END -->`, ""} <div class="text-center"><div class="max-w-3xl mx-auto"><h1 class="text-4xl font-bold text-gray-900 mb-4" data-svelte-h="svelte-dlw63r">Link shortening <span class="text-eco-400">Sustentável</span></h1> <p class="text-xl text-gray-900 mb-8" data-svelte-h="svelte-127dpbm">Turn long URLs into short, eco-friendly links with included QR Code</p> <div class="bg-white rounded-xl shadow-lg p-8 mb-8"><form class="space-y-4"><div><input type="url" placeholder="Paste your URL here..." class="input-eco w-full text-lg" required${add_attribute("value", url, 0)}></div> <button type="submit" ${""} class="btn-eco text-lg px-8 py-3 disabled:opacity-50">${escape("🌱 Shorten Link")}</button></form></div> <div class="grid md:grid-cols-3 gap-8 text-center" data-svelte-h="svelte-mklzek"><div class="bg-white p-6 rounded-lg shadow"><div class="text-3xl mb-2">⚡</div> <h3 class="font-semibold mb-2">Ultra Fast</h3> <p class="text-gray-600">Redirection in milliseconds</p></div> <div class="bg-white p-6 rounded-lg shadow"><div class="text-3xl mb-2">🌱</div> <h3 class="font-semibold mb-2">Sustainable</h3> <p class="text-gray-600">Eco-friendly Infraestructure</p></div> <div class="bg-white p-6 rounded-lg shadow"><div class="text-3xl mb-2">📊</div> <h3 class="font-semibold mb-2">Analytics</h3> <p class="text-gray-600">Detailed metrics</p></div></div></div></div>`;
});
export {
  Page as default
};
