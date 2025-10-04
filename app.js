// Telegram Web App инициализация
const tg = window.Telegram.WebApp;

// Инициализация приложения
function initApp() {
    // Развернуть на весь экран
    tg.expand();
    
    // Показать основную кнопку
    tg.MainButton.setText("Обновить данные");
    tg.MainButton.show();
    
    // Обработчик основной кнопки
    tg.MainButton.onClick(() => {
        tg.showPopup({
            title: "Обновление",
            message: "Данные успешно обновлены!",
            buttons: [{ type: "ok" }]
        });
    });
    
    // Получить данные пользователя
    const user = tg.initDataUnsafe.user;
    displayUserInfo(user);
    
    console.log("Math Tutor Pro запущен!", user);
}

// Отображение информации о пользователе
function displayUserInfo(user) {
    const userInfoEl = document.getElementById('user-info');
    
    if (user) {
        userInfoEl.innerHTML = `
            <h3>👋 Привет, ${user.first_name}!</h3>
            <p><strong>ID:</strong> ${user.id}</p>
            ${user.username ? `<p><strong>Username:</strong> @${user.username}</p>` : ''}
            <p><strong>Статус:</strong> Репетитор</p>
        `;
    } else {
        userInfoEl.innerHTML = `
            <h3>Демо-режим</h3>
            <p>Приложение запущено вне Telegram</p>
        `;
    }
}

// Запуск приложения при загрузке
document.addEventListener('DOMContentLoaded', initApp);