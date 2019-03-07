import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthenticationProvider } from 'app/model/user.model';
import { AuthentificationStore, UserService } from 'app/service/services.module';
import { AccountComponent } from '../account.component';

@Component({
    selector: 'app-account-auth',
    templateUrl: './auth.html',
    styleUrls: ['./auth.scss']
})

export class AuthComponent extends AccountComponent implements OnInit {
    authProviders: Array<AuthenticationProvider>;

    ngOnInit(): void {
        this._userService.loginRedirect().subscribe((response) => {
            this.authProviders = response;
        });
    }

    constructor(private _userService: UserService, // private _router: Router,
        _authStore: AuthentificationStore, private _route: ActivatedRoute) {
        super(_authStore);

        this._route.queryParams.subscribe(queryParams => {
            console.log(queryParams);
        });
    }

    submitForm(form: any): void {
        console.log(form);
    }

}
