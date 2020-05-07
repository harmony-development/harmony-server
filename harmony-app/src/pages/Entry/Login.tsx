import React, { useRef, useEffect } from 'react';
import { TextField, Typography, Button, makeStyles, Theme } from '@material-ui/core';
import { toast } from 'react-toastify';
import { useHistory } from 'react-router';
import { useTranslation } from 'react-i18next';

import { harmonySocket } from '../../Root';

const loginStyles = makeStyles((theme: Theme) => ({
	root: {
		paddingLeft: theme.spacing(1),
		paddingRight: theme.spacing(1),
		paddingTop: theme.spacing(1),
		paddingBottom: theme.spacing(1),
	},
	submitBtn: {
		marginTop: theme.spacing(2),
	},
}));

export const Login = () => {
	const history = useHistory();
	const classes = loginStyles();
	const { t } = useTranslation('entry');

	const [err, setErr] = React.useState<string | undefined>(undefined);
	const emailRef = useRef<HTMLInputElement | undefined>(undefined);
	const pwdRef = useRef<HTMLInputElement | undefined>(undefined);

	const login = () => {
		if (harmonySocket.conn.readyState === WebSocket.CLOSED) {
			toast.error("Can't login, trouble connecting to server");
		} else if (emailRef.current && pwdRef.current && emailRef.current.value.length && pwdRef.current.value) {
			harmonySocket.login(emailRef.current.value, pwdRef.current.value);
		} else {
			toast.error("Can't login, missing email or password");
		}
	};

	useEffect(() => {
		harmonySocket.events.addListener('loginerror', (raw: any) => {
			if (typeof raw['message'] === 'string') {
				setErr(raw['message']);
			}
		});
		harmonySocket.events.addListener('token', (raw: any) => {
			if (typeof raw['token'] === 'string' && typeof raw['userid'] === 'string') {
				localStorage.setItem('token', raw['token']);
				localStorage.setItem('userid', raw['userid']);
				harmonySocket.refreshToken();
				harmonySocket.getGuilds();
				harmonySocket.getSelf();
				history.push('/app');
			}
		});
		return () => {
			harmonySocket.events.removeAllListeners('loginerror');
			harmonySocket.events.removeAllListeners('token');
		};
	}, [history]);

	return (
		<div className={classes.root}>
			<form onSubmit={(e: React.FormEvent<EventTarget>) => e.preventDefault()}>
				<TextField
					label={t('entry:email')}
					type="email"
					name="email"
					autoComplete="email"
					margin="normal"
					fullWidth
					inputRef={emailRef}
				/>
				<TextField
					label={t('entry:password')}
					type="password"
					name="password"
					margin="normal"
					fullWidth
					inputRef={pwdRef}
				/>
				{err ? (
					<Typography variant="subtitle1" color={'error'}>
						{err}
					</Typography>
				) : (
					undefined
				)}
				<Button
					variant="contained"
					color="secondary"
					className={classes.submitBtn}
					onClick={login}
					type="submit"
					fullWidth
				>
					{t('entry:log-in')}
				</Button>
			</form>
		</div>
	);
};