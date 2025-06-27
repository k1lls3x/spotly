// src/App.tsx (иконки убраны, инпут подогнан под рамки)
import { useEffect, useState } from 'react';
import './App.css';
import Logo from './components/Logo';

const cities = [
  'Ростов-на-Дону',
  'Москва',
  'Санкт-Петербург',
  'Сочи',
  'Краснодар'
];

const categories = [
  'Клубы',
  'Достопримечательности',
  'Театры',
  'Музеи',
  'Мероприятия',
  'Рестораны',
  'Кафе'
];

export default function App() {
  const [search, setSearch] = useState('');
  const [selectedCity, setSelectedCity] = useState<string | null>(null);
  const [confirmed, setConfirmed] = useState(false);
  const [showCategories, setShowCategories] = useState(false);

  const filteredCities = cities.filter((city) =>
    city.toLowerCase().includes(search.toLowerCase())
  );

  useEffect(() => {
    if (confirmed) {
      const timer = setTimeout(() => setShowCategories(true), 500);
      return () => clearTimeout(timer);
    }
  }, [confirmed]);

  return (
    <div className="app-wrapper">
      <Logo />
      <div className="city-box">
        <h2 className="city-title">
          {selectedCity ? `Город: ${selectedCity}` : 'Выбрать город'}
        </h2>

        {!confirmed && (
          <>
            <input
              className="city-search"
              type="text"
              placeholder="Введите город..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
            <div className="city-options">
              {filteredCities.map((city) => (
                <div
                  key={city}
                  className="city-option"
                  onClick={() => setSelectedCity(city)}
                >
                  {city}
                </div>
              ))}
            </div>
            {selectedCity && (
              <button className="select-btn gradient" onClick={() => setConfirmed(true)}>
                Выбрать
              </button>
            )}
          </>
        )}

        {confirmed && showCategories && (
          <div className="category-grid">
            {categories.map((label, index) => (
              <div
                key={label}
                className="category-btn smooth"
                style={{ animationDelay: `${index * 100}ms` }}
              >
                {label}
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
