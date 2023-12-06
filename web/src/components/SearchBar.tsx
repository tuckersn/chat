import { FC, useState } from 'preact/compat';
import { SiAddthis, SiElasticsearch, SiOpensearch, SiZincsearch } from 'react-icons/si';
import { AiOutlineSearch } from 'react-icons/ai';

export const SearchBar: FC<{
	onSearch: (searchTerm: string) => void;
}> = ({ onSearch }) => {
	const [searchTerm, setSearchTerm] = useState('');

	const handleSearch = () => {
		onSearch(searchTerm);
	};

	return (
		<div className="search-bar">
			<input 
				type="text" 
				className="search-input" 
				value={searchTerm} 
				onChange={(e) => setSearchTerm((e as any).target.value)} 
				placeholder="Search..."
			/>
			<button className="search-button" onClick={handleSearch}>
				<AiOutlineSearch/>
			</button>
			<style jsx>{`
				.search-bar {
					display: flex;
					border: 1px solid #ccc;
					border-radius: 4px;
					overflow: hidden;
				}
				.search-input {
					flex: 1;
					padding: 8px;
					border: none;
					outline: none;
				}
				.search-button {
					padding: 8px 15px;
					border: none;
					background-color: #007bff;
					color: white;
					cursor: pointer;
					outline: none;
				}
				.search-button:hover {
					background-color: #0056b3;
				}
			`}</style>
		</div>
	);
};