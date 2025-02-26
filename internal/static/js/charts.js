// Добавляем взаимодействие с фильтрами
document.querySelectorAll('.filter-btn').forEach(btn => {
    btn.addEventListener('click', function() {
        document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
        this.classList.add('active');
    });
});

// Обработка вкладок жанров
document.querySelectorAll('.genre-tab').forEach(tab => {
    tab.addEventListener('click', function() {
        document.querySelectorAll('.genre-tab').forEach(t => t.classList.remove('active'));
        this.classList.add('active');
    });
});