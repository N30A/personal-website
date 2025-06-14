(function () {
  try {
    if (localStorage.getItem("theme") === "dark") {
      document.documentElement.classList.add("dark");
    }
  } catch (e) { }
})();

document.addEventListener("DOMContentLoaded", () => {
  const toggleButton = document.getElementById("theme-toggle");

  if (document.documentElement.classList.contains("dark")) {
    toggleButton.textContent = "🌑";
  } else {
    toggleButton.textContent = "🌕";
  }

  toggleButton.addEventListener("click", () => {
    document.documentElement.classList.toggle("dark");
    if (document.documentElement.classList.contains("dark")) {
      localStorage.setItem("theme", "dark");
      toggleButton.textContent = "🌑"
    } else {
      localStorage.setItem("theme", "light");
      toggleButton.textContent = "🌕"
    }
  });
});
