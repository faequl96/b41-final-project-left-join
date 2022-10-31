const nama = document.querySelector('.nama');
const email = document.querySelector('.email');
const phone = document.querySelector('.phone');
const subject = document.querySelector('.subject');
const message = document.querySelector('.message');


nama.addEventListener('click', function() {
    
    nama.classList.add('pin');
    nama.removeAttribute('placeholder');
    if(email.classList.contains('pin') && email.value == "") {
        email.placeholder = '* Email wajib diinput';
        email.style.border = '2px solid red';
    }
    if(phone.classList.contains('pin') && phone.value == "") {
        phone.placeholder = '* Phone wajib diinput';
        phone.style.border = '2px solid red';
    }
    if(message.classList.contains('pin') && message.value == "") {
        message.placeholder = '* Message wajib diinput';
        message.style.border = '2px solid red';
    }
    if(subject.classList.contains('pin')) {
        subject.style.border = '2px solid red';
    }
    nama.addEventListener('input', function() {
        if(nama.value == "") {
            nama.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
        } else{
            nama.style.border = '2px solid rgba(0,0,0,0)';
            nama.classList.remove('pin');
            if(nama.value != "" && email.value != "" && phone.value != "" && subject.value != "" && message.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
            }
        }
    })
})


email.addEventListener('click', function() {
    
    email.classList.add('pin');
    email.removeAttribute('placeholder');
    if(nama.classList.contains('pin') && nama.value == "") {
        nama.placeholder = '* Nama wajib diinput';
        nama.style.border = '2px solid red';
    }
    if(phone.classList.contains('pin') && phone.value == "") {
        phone.placeholder = '* Phone wajib diinput';
        phone.style.border = '2px solid red';
    }
    if(message.classList.contains('pin') && message.value == "") {
        message.placeholder = '* Message wajib diinput';
        message.style.border = '2px solid red';
    }
    if(subject.classList.contains('pin')) {
        subject.style.border = '2px solid red';
    }
    email.addEventListener('input', function() {
        if(email.value == "") {
            email.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
        } else{
            email.style.border = '2px solid rgba(0,0,0,0)';
            email.classList.remove('pin');
            if(nama.value != "" && email.value != "" && phone.value != "" && subject.value != "" && message.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
            }
        }
    })
})


phone.addEventListener('click', function() {
    
    phone.classList.add('pin');
    phone.removeAttribute('placeholder');
    if(nama.classList.contains('pin') && nama.value == "") {
        nama.placeholder = '* Nama wajib diinput';
        nama.style.border = '2px solid red';
    }
    if(email.classList.contains('pin') && email.value == "") {
        email.placeholder = '* Email wajib diinput';
        email.style.border = '2px solid red';
    }
    if(message.classList.contains('pin') && message.value == "") {
        message.placeholder = '* Message wajib diinput';
        message.style.border = '2px solid red';
    }
    if(subject.classList.contains('pin')) {
        subject.style.border = '2px solid red';
    }
    phone.addEventListener('input', function() {
        if(phone.value == "") {
            phone.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
        } else{
            phone.style.border = '2px solid rgba(0,0,0,0)';
            phone.classList.remove('pin');
            if(nama.value != "" && email.value != "" && phone.value != "" && subject.value != "" && message.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
            }
        }
    })
})


subject.addEventListener('click', function() {
    
    if(nama.classList.contains('pin') && nama.value == "") {
        nama.placeholder = '* Nama wajib diinput';
        nama.style.border = '2px solid red';
    }
    if(phone.classList.contains('pin') && phone.value == "") {
        phone.placeholder = '* Phone wajib diinput';
        phone.style.border = '2px solid red';
    }
    if(email.classList.contains('pin') && email.value == "") {
        email.placeholder = '* Email wajib diinput';
        email.style.border = '2px solid red';
    }
    if(message.classList.contains('pin') && message.value == "") {
        message.placeholder = '* Message wajib diinput';
        message.style.border = '2px solid red';
    }
    if(subject.value == "") {
        subject.classList.add('pin');
    }
    
    subject.addEventListener('input', function() {
        if(subject.value == "") {
            subject.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
        } else{
            subject.style.border = '2px solid rgba(0,0,0,0)'
            subject.classList.remove('pin');
            if(nama.value != "" && email.value != "" && phone.value != "" && subject.value != "" && message.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
            }
        }
    })
})


message.addEventListener('click', function() {
    
    message.classList.add('pin');
    message.removeAttribute('placeholder');
    if(nama.classList.contains('pin') && nama.value == "") {
        nama.placeholder = '* Nama wajib diinput';
        nama.style.border = '2px solid red';
    }
    if(phone.classList.contains('pin') && phone.value == "") {
        phone.placeholder = '* Phone wajib diinput';
        phone.style.border = '2px solid red';
    }
    if(email.classList.contains('pin') && email.value == "") {
        email.placeholder = '* Email wajib diinput';
        email.style.border = '2px solid red';
    }
    if(subject.classList.contains('pin') && subject.value == "") {
        subject.style.border = '2px solid red';
    }
    message.addEventListener('input', function() {
        if(message.value == "") {
            message.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
        } else{
            message.style.border = '2px solid rgba(0,0,0,0)';
            message.classList.remove('pin');
            if(nama.value != "" && email.value != "" && phone.value != "" && subject.value != "" && message.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
            }
        }
    })
})
