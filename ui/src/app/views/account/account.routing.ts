import { ModuleWithProviders } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';
import { LoginComponent } from './login/login.component';
import { PasswordComponent } from './password/password.component';
import { SignUpComponent } from './signup/signup.component';
import { VerifyComponent } from './verify/verify.component';

const routes: Routes = [
    {
        path: '',
        children : [
            { path: '', redirectTo: 'auth', pathMatch: 'full' },
            {
                path: 'auth',
                component: AuthComponent,
                data: { title: 'CDS • Authentication' }
            },
            {
                path: 'login',
                component: LoginComponent,
                data: { title: 'CDS • Login' }
            },
            { path: 'password', component: PasswordComponent, data: { title: 'Reset Password' }},
            { path: 'signup', component: SignUpComponent, data: { title: 'Signup' }},
            { path: 'verify/:username/:token', component: VerifyComponent }
        ]
    }
];

export const accountRouting: ModuleWithProviders = RouterModule.forChild(routes);
