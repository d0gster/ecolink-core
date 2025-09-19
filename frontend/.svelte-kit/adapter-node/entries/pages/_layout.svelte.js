import { c as create_ssr_component, a as add_attribute, b as subscribe, e as escape, o as onDestroy, v as validate_component } from "../../chunks/ssr.js";
import { i as isAuthenticated, u as user, a as isLoading } from "../../chunks/auth.js";
import { d as derived } from "../../chunks/index.js";
const app = "";
const BackgroundVideo_svelte_svelte_type_style_lang = "";
const css = {
  code: "video.svelte-65inv1{filter:saturate(1.2) brightness(0.9)}",
  map: null
};
const BackgroundVideo = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { desktopSrc = "" } = $$props;
  let { mobileSrc = "" } = $$props;
  let { poster = "" } = $$props;
  let videoElement;
  let isMobile = false;
  if ($$props.desktopSrc === void 0 && $$bindings.desktopSrc && desktopSrc !== void 0)
    $$bindings.desktopSrc(desktopSrc);
  if ($$props.mobileSrc === void 0 && $$bindings.mobileSrc && mobileSrc !== void 0)
    $$bindings.mobileSrc(mobileSrc);
  if ($$props.poster === void 0 && $$bindings.poster && poster !== void 0)
    $$bindings.poster(poster);
  $$result.css.add(css);
  {
    if (typeof window !== "undefined") {
      isMobile = window.innerWidth < 768;
    }
  }
  return `<div class="fixed inset-0 -z-10 overflow-hidden"><video class="absolute inset-0 w-full h-full object-cover opacity-70 svelte-65inv1" autoplay muted loop playsinline${add_attribute("poster", poster, 0)}${add_attribute("this", videoElement, 0)}><source${add_attribute("src", isMobile ? mobileSrc : desktopSrc, 0)} type="video/mp4"><div class="absolute inset-0 bg-gradient-to-br from-green-50 to-green-100"></div></video>  <div class="absolute inset-0 bg-white/20 backdrop-blur-[0.5px]"></div> </div>`;
});
const UserDropdown = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $isAuthenticated, $$unsubscribe_isAuthenticated;
  let $user, $$unsubscribe_user;
  $$unsubscribe_isAuthenticated = subscribe(isAuthenticated, (value) => $isAuthenticated = value);
  $$unsubscribe_user = subscribe(user, (value) => $user = value);
  $$unsubscribe_isAuthenticated();
  $$unsubscribe_user();
  return `${$isAuthenticated && $user ? `<div class="relative"><button class="flex items-center space-x-2 text-gray-700 hover:text-eco-600 focus:outline-none"><img${add_attribute("src", $user.picture || "https://ui-avatars.com/api/?name=" + encodeURIComponent($user.name || $user.email || "Usuario") + "&background=16a34a&color=fff", 0)}${add_attribute("alt", $user.name || "Usuario", 0)} class="w-8 h-8 rounded-full" referrerpolicy="no-referrer" crossorigin="anonymous"> <span class="hidden md:block font-medium">${escape($user.name)}</span> <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg></button> ${``}</div>` : `<a href="/auth" class="text-gray-700 hover:text-eco-600 font-medium" data-svelte-h="svelte-1y5io9w">Login</a>`}  `;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $isLoading, $$unsubscribe_isLoading;
  $$unsubscribe_isLoading = subscribe(isLoading, (value) => $isLoading = value);
  const showSkeleton = derived(isLoading, ($isLoading2) => $isLoading2 === true);
  let skeletonUnsub = showSkeleton.subscribe(() => {
  });
  onDestroy(() => skeletonUnsub());
  $$unsubscribe_isLoading();
  return `<div class="min-h-screen relative"> ${validate_component(BackgroundVideo, "BackgroundVideo").$$render(
    $$result,
    {
      desktopSrc: "/videos/eco-background-desktop.mp4",
      mobileSrc: "/videos/eco-background-mobile.mp4",
      poster: ""
    },
    {},
    {}
  )} <nav class="bg-white shadow-sm border-b border-eco-200"><div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"><div class="flex justify-between h-16"><div class="flex items-center" data-svelte-h="svelte-14mxmy9"><h1 class="text-2xl font-bold text-eco-700">🌱 EcoLink</h1></div> <div class="flex items-center space-x-4"><a href="/" class="text-gray-700 hover:text-eco-600" data-svelte-h="svelte-hzhxvq">Home</a> <a href="/dashboard" class="text-gray-700 hover:text-eco-600" data-svelte-h="svelte-fiu9kn">Dashboard</a> ${validate_component(UserDropdown, "UserDropdown").$$render($$result, {}, {}, {})}</div></div></div></nav> ${$isLoading ? ` <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8" data-svelte-h="svelte-5ixntd"><div class="animate-pulse space-y-4"><div class="h-6 bg-gray-200 rounded w-1/3"></div> <div class="h-4 bg-gray-200 rounded w-full"></div> <div class="h-48 bg-gray-200 rounded w-full"></div></div></main>` : `<main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">${slots.default ? slots.default({}) : ``}</main>`}</div>`;
});
export {
  Layout as default
};
