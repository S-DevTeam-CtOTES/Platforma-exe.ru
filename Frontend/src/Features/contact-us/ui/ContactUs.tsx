import React from 'react'
import { App } from "@Shared/consts/index";
import ImageContactUs from '@Shared/assets/img/ImageContactUs.png';
import './ContactUs.scss';

const ContactUs = () => {
    return (
        <div className='contactUs'>
            <h2 className='contactUs__Title'>Свяжитесь с нами</h2>
            <form className='contactUs__form form' action="">
                <img className='form__image' src={ImageContactUs} alt="" />
                <div className='form__wrapper'>
                    <div className='form__container'>
                        <div className='form__block'>
                            <div className='form__input-block'>
                                <input className='form__input' type="text" placeholder='Полное имя' />
                            </div>
                            <div className='form__input-block'>
                                <input className='form__input' type="text" placeholder='Компания' />
                            </div>
                        </div>
                        <div className='form__block'>
                            <div className='form__input-block'>
                                <input className='form__input' type="text" placeholder='Адрекс электронной почты' />
                            </div>
                            <div className='form__input-block'>
                                <input className='form__input' type="text" placeholder='Адрес веб-сайта' />
                            </div>
                        </div>
                    </div>
                    <div className='form__textarea-block'>
                        <textarea className='form__textarea' placeholder='Ваше сообщение' name=""></textarea>
                    </div>

                </div>
                <button className='form__button' type='submit'>Отправить</button>
            </form>
            <div className='contactUs__copyright'>{App.Copyright}</div>
        </div>
    )
}

export { ContactUs }