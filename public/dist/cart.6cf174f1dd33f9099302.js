!function(){if(console.log("cart.js"),document.getElementById&&document.createElement){const e=document.getElementById("cart-form"),t=document.getElementById("cart-btn"),n=document.createElement("div");n.className="cart-menu";const c=document.createElement("p");if(c.textContent="Cart menu",n.append(c),t.addEventListener){let c=!1;t.addEventListener("click",(()=>{c?e.removeChild(n):e.append(n),c=!c}))}}}();