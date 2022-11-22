(function() {
  console.log('cart.js');
  if (document.getElementById && document.createElement) {
    const cartForm = document.getElementById('cart-form');
    const cartBtn = document.getElementById('cart-btn');
    const el = document.createElement('div');
    el.className = "cart-menu";
    const p = document.createElement('p');
    p.textContent = 'Cart menu';
    el.append(p);

    if (cartBtn.addEventListener) {
      let open = false;
      cartBtn.addEventListener('click', () => {
        if (open) {
          cartForm.removeChild(el);
        } else {
          cartForm.append(el);
        }
        open = !open;
      });
    }
  }
})()