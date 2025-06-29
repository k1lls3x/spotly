import { useState } from 'react';
import './App.css';
import Logo from './components/Logo';

const cities = ['Ростов-на-Дону', 'Москва', 'Санкт-Петербург', 'Сочи'];
const categories = ['Кафе', 'Рестораны', 'Клубы', 'Достопримечательности', 'Мероприятия'];

// Пример данных для анкет (можно заменить на реальные)
const cards = [
  {
    title: 'Кафе Латте',
    description: 'Уютное кафе с отличным кофе и атмосферой для встреч.',
    images: [
      'https://placehold.co/400x260?text=Latte1',
      'https://placehold.co/400x260?text=Latte2',
      'https://placehold.co/400x260?text=Latte3',
    ],
  },
  {
    title: 'Ресторан Сити',
    description: 'Современный ресторан с авторской кухней и красивым видом.',
    images: [
      'https://placehold.co/400x260?text=Rest1',
      'https://placehold.co/400x260?text=Rest2',
    ],
  },
  {
    title: 'Ночной клуб Neon',
    description: 'Лучшее место для весёлых вечеринок и новых знакомств.',
    images: [
      'https://placehold.co/400x260?text=Club1',
      'https://placehold.co/400x260?text=Club2',
    ],
  }
];

function App() {
  const [selectedCity, setSelectedCity] = useState<string>('');
  const [cityInput, setCityInput] = useState<string>('');
  const [showSelectButton, setShowSelectButton] = useState(false);
  const [confirmedCity, setConfirmedCity] = useState<string | null>(null);
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null);

  // Для анкет
  const [cardIndex, setCardIndex] = useState(0);
  const [photoIndex, setPhotoIndex] = useState(0);
  const [fullscreen, setFullscreen] = useState(false);

  const filteredCities = cities.filter((city) =>
    city.toLowerCase().includes(cityInput.toLowerCase())
  );

  const handleConfirmCity = () => {
    if (selectedCity) {
      setConfirmedCity(selectedCity);
    }
  };

  const handleCategoryClick = (category: string) => {
    setSelectedCategory(category);
  };

  return (
    <div className="app-wrapper">
      <Logo />

      {/* Выбор города */}
      {!confirmedCity && (
        <div className="city-box">
          <h2 className="city-title">Выбрать город</h2>
          <input
            className="city-search"
            placeholder="Введите город..."
            value={cityInput}
            onChange={(e) => setCityInput(e.target.value)}
            onFocus={() => setShowSelectButton(true)}
          />

          <div className="city-options">
            {filteredCities.map((city) => (
              <div
                key={city}
                className={`city-option ${selectedCity === city ? 'selected' : ''}`}
                onClick={() => setSelectedCity(city)}
              >
                {city}
              </div>
            ))}
          </div>

          {selectedCity && (
            <button className="select-btn gradient" onClick={handleConfirmCity}>
              Выбрать
            </button>
          )}
        </div>
      )}

      {/* После выбора города — отображаем выбранный город и категории */}
      {confirmedCity && !selectedCategory && (
        <>
          <div className="city-box">
            <h2 className="city-title">Выбранный город: {confirmedCity}</h2>
          </div>
          <div className="category-grid">
            {categories.map((category) => (
              <div
                key={category}
                className="category-btn"
                onClick={() => handleCategoryClick(category)}
              >
                {category}
              </div>
            ))}
          </div>
        </>
      )}

      {/* Анкета */}
      {confirmedCity && selectedCategory && (
        <div className="card-view">
          {/* Стрелки для листания анкет */}
          <button
            className="arrow-btn card-left"
            onClick={() => {
              setCardIndex((prev) => prev === 0 ? cards.length - 1 : prev - 1);
              setPhotoIndex(0);
            }}
            tabIndex={-1}
            aria-label="Предыдущая анкета"
          >
            &#8592;
          </button>

          {/* Фото анкеты */}
          <div
            className="card-photo"
            onClick={() => setFullscreen(true)}
            style={{
              backgroundImage: `url(${cards[cardIndex].images[photoIndex]})`
            }}
            title="Открыть на весь экран"
          ></div>

          {/* Стрелка вправо */}
          <button
            className="arrow-btn card-right"
            onClick={() => {
              setCardIndex((prev) => prev === cards.length - 1 ? 0 : prev + 1);
              setPhotoIndex(0);
            }}
            tabIndex={-1}
            aria-label="Следующая анкета"
          >
            &#8594;
          </button>

          {/* Индикаторы фото */}
          <div className="card-indicators">
            {cards[cardIndex].images.map((_, idx) => (
              <span
                key={idx}
                className={`dot${photoIndex === idx ? ' active' : ''}`}
                onClick={() => setPhotoIndex(idx)}
              />
            ))}
          </div>
          {/* Название и описание */}
          <div className="card-title" style={{ fontWeight: 600, marginBottom: 8, color: 'var(--text-color)' }}>
            {cards[cardIndex].title}
          </div>
          <div className="card-text">
            {cards[cardIndex].description}
          </div>

          {/* ====== Полноэкранное фото ====== */}
          {fullscreen && (
            <div className="modal-overlay" onClick={() => setFullscreen(false)}>
              <button
                className="modal-arrow"
                style={{ left: 0, position: 'absolute' }}
                onClick={e => { e.stopPropagation(); setPhotoIndex((prev) => prev === 0 ? cards[cardIndex].images.length - 1 : prev - 1); }}
              >&#8592;</button>
              <img
                src={cards[cardIndex].images[photoIndex]}
                alt="Фото"
                className="modal-photo"
                onClick={e => e.stopPropagation()}
              />
              <button
                className="modal-arrow"
                style={{ right: 0, position: 'absolute' }}
                onClick={e => { e.stopPropagation(); setPhotoIndex((prev) => prev === cards[cardIndex].images.length - 1 ? 0 : prev + 1); }}
              >&#8594;</button>
              <button
                className="modal-close"
                onClick={e => { e.stopPropagation(); setFullscreen(false); }}
              >✕</button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default App;
