const email = document.querySelector('.email');
const password = document.querySelector('.password');
const button = document.querySelector('.submit button');


email.addEventListener('click', function() {
    
    email.classList.add('pin');
    email.removeAttribute('placeholder');
    if(password.classList.contains('pin') && password.value == "") {
        password.placeholder = '* Password wajib diinput';
        password.style.border = '2px solid red';
    }
    email.addEventListener('input', function() {
        if(email.value == "") {
            email.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
            button.removeAttribute('type');
            button.setAttribute('type', 'button');
        } else{
            email.style.border = '2px solid rgba(0,0,0,0)';
            email.classList.remove('pin');
            if(email.value != "" && password.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
                button.removeAttribute('type');
                button.setAttribute('type', 'submit');
            }
        }
    })
})


password.addEventListener('click', function() {
    
    password.classList.add('pin');
    password.removeAttribute('placeholder');
    if(email.classList.contains('pin') && email.value == "") {
        email.placeholder = '* Email wajib diinput';
        email.style.border = '2px solid red';
    }
    password.addEventListener('input', function() {
        if(password.value == "") {
            password.style.border = '2px solid red';
            button.style.backgroundColor = '#e4e4e4';
            button.style.color = '#b2abab';
            button.style.cursor = 'default';
            button.removeAttribute('type');
            button.setAttribute('type', 'button');
        } else{
            password.style.border = '2px solid rgba(0,0,0,0)';
            password.classList.remove('pin');
            if(email.value != "" && password.value != "") {
                button.style.backgroundColor = '#000';
                button.style.color = '#eaeaea';
                button.style.cursor = 'pointer';
                button.removeAttribute('type');
                button.setAttribute('type', 'submit');
            }
        }
    })
})


button.addEventListener('click', function() {

    if(email.value != "" && password.value != "") {
        
    } else {
        alert('Semua field wajib diisi');
    }
})