(function() {
  if (document.getElementById) {
    const cartForm = document.getElementById('cart-form');
    const cartBtn = document.getElementById('cart-btn');
    
    if (cartForm.addEventListener && cartBtn.addEventListener) {
      cartForm.addEventListener('submit', (e) => {
        e.preventDefault();
      })
      cartBtn.addEventListener('focus', (e) => {
        e.preventDefault();
        import('./cart');
      })
    }
  }
})()