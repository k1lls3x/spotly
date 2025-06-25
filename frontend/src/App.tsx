// src/App.tsx
import { useState } from 'react';
import './App.css';
import CardList from './components/CardList';
import Logo from './components/Logo';

const categories = [
  'Музеи',
  'Достопримечательности',
  'Клубы',
  'Мероприятия',
  'Рестораны',
  'Кафе'
];

function App() {
  const [selectedCity, setSelectedCity] = useState('Ростов-на-Дону');
  const [selectedCategory, setSelectedCategory] = useState<string | null>(null);

  const cards: any[] = []; // временно пусто, будут приходить из бэка позже

  return (
    <div className="app">
      <header className="app-header">
        <Logo />
        <h1 className="app-title">Spotly</h1>
        <div className="city-select">
          <span>Город: </span>
          <select
            value={selectedCity}
            onChange={(e) => setSelectedCity(e.target.value)}
          >
            <option>Ростов-на-Дону</option>
            <option>Москва</option>
            <option>Санкт-Петербург</option>
            <option>Сочи</option>
            <option>Краснодар</option>
          </select>
        </div>
      </header>

      <div className="category-list">
        {categories.map((cat) => (
          <button
            key={cat}
            className={`category-btn ${selectedCategory === cat ? 'active' : ''}`}
            onClick={() => setSelectedCategory(cat)}
          >
            {cat}
          </button>
        ))}
      </div>

      <div className="content">
        {!selectedCategory && <p>Выберите категорию выше</p>}
        {selectedCategory && <CardList cards={cards} />}
      </div>
    </div>
  );
}

export default App;

// самая главная страница ..