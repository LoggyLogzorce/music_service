// Обновлённый скрипт с медленной прокруткой
let currentIndex = 0;
const slides = document.querySelectorAll('.carousel-item');
const totalSlides = slides.length;

function showSlide(index) {
    const carouselInner = document.querySelector('.carousel-inner');
    const offset = -index * 100;
    carouselInner.style.transform = `translateX(${offset}%)`;
}

function nextSlide() {
    currentIndex = (currentIndex + 1) % totalSlides;
    showSlide(currentIndex);
}

function prevSlide() {
    currentIndex = (currentIndex - 1 + totalSlides) % totalSlides;
    showSlide(currentIndex);
}

// Замедляем автоматическую прокрутку до 7 секунд
setInterval(nextSlide, 7000);

// Добавляем плавность анимации
document.querySelector('.carousel-inner').style.transition = 'transform 0.8s ease-in-out';