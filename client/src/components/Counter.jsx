import { lazy, Suspense } from 'preact/compat';

const Message = lazy(async () => import('./Message'));
const Fallback = () => <p>Loading...</p>;

export default function Counter({ children, count }) {
	const add = () => count.value++;
	const subtract = () => count.value--;

	return (
		<>
			<div>
				<button onClick={subtract}>-</button>
				<pre>{count}</pre>
				<button onClick={add}>+</button>
			</div>
			<Suspense fallback={Fallback}>
				<Message>{children}</Message>
			</Suspense>
		</>
	);
}
