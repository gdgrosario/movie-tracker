import { Directive } from '@angular/core';
import {
  NG_VALIDATORS,
  Validator,
  AbstractControl,
  ValidationErrors,
  ValidatorFn
} from '@angular/forms';

@Directive({
  selector: '[appPasswordMatch]',
  providers: [{
    provide: NG_VALIDATORS,
    useExisting: PasswordMatchValidatorDirective,
    multi: true
  }]
})
export class PasswordMatchValidatorDirective implements Validator {
  validate(control: AbstractControl): ValidationErrors | null {
    return passwordMatchValidator(control);
  }
}

export const passwordMatchValidator: ValidatorFn = (control: AbstractControl): ValidationErrors | null => {
  const password = control.get('password');
  const repeatPassword = control.get('repeatPassword');

  const error = password && repeatPassword && password.value !== repeatPassword.value ? { passwordsNotMatch: true } : null;
  repeatPassword?.setErrors(error);

  return error;
};
