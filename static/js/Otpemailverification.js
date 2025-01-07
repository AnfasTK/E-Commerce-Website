function moveFocus(current, nextFieldId) {
    if (current.value.length === 1 && nextFieldId) {
      document.getElementById(nextFieldId).focus();
    }
  }

  let countdown = 30;
  const timerElement = document.getElementById('timer');
  const resendButton = document.getElementById('resendButton');

  const timerInterval = setInterval(() => {
    countdown--;
    if (countdown > 0) {
      timerElement.textContent = `Resend OTP in ${countdown} seconds`;
    } else {
      clearInterval(timerInterval);
      timerElement.classList.add('hidden');
      resendButton.classList.remove('hidden');
    }
  }, 1000);

  function resendOTP() {
    // Simulate OTP resend logic
    countdown = 30;
    timerElement.textContent = `Resend OTP in ${countdown} seconds`;
    timerElement.classList.remove('hidden');
    resendButton.classList.add('hidden');
    setInterval(() => {
      countdown--;
      if (countdown > 0) {
        timerElement.textContent = `Resend OTP in ${countdown} seconds`;
      } else {
        clearInterval(timerInterval);
        timerElement.classList.add('hidden');
        resendButton.classList.remove('hidden');
      }
    }, 1000);
  }