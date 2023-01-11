(function() {
  async function registerCachingServiceWorker() {
    try {
      await navigator.serviceWorker.register('/caching-service-worker.js');
      console.log('Caching service worker registered!');
    } catch (err) {
      console.error(`Registration failed: ${err}`);
    }
  }

  function dynamicallyImportCart() {
    if (!document.getElementById) {
      return;
    }
    const cartForm = document.getElementById('cart-form');
    const cartBtn = document.getElementById('cart-btn');

    if (!cartForm || !cartBtn) {
      return;
    }
    if (!cartForm.addEventListener || !cartBtn.addEventListener) {
      return;
    }

    cartForm.addEventListener('submit', (e) => {
      e.preventDefault();
    });
    cartBtn.addEventListener('focus', (e) => {
      e.preventDefault();
      import('./cart.mjs');
    });
  }

  dynamicallyImportCart();
  if (navigator.serviceWorker) {
    registerCachingServiceWorker();
  }
})();